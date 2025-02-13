package middleware

import (
	"log"
	"techtest/configs"
	"techtest/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal register",
		})
	}

	// Simpan user ke database MySQL dengan GORM
	user.Password = string(hashedPassword)
	if err := configs.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal register",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil register!",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parsing body request ke struct loginData
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// Cari user di database berdasarkan username
	var user models.User
	if err := configs.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Username atau password salah",
			})
		}
		log.Println("Error saat mencari user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal login",
		})
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// Buat JWT token
	claims := jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.JwtSecret)) // Menggunakan secret dari configs
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func GetMe(c *fiber.Ctx) error {
	// Mendapatkan data user yang sedang login melalui JWT token
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token tidak valid",
		})
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Gagal membaca klaim token",
		})
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User ID tidak valid",
		})
	}

	// Cari user di database berdasarkan user ID
	var userData models.User
	if err := configs.DB.First(&userData, uint(userID)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan data user",
		})
	}

	return c.JSON(fiber.Map{
		"user": userData,
	})
}

// Middleware untuk otentikasi JWT
func Authenticate(c *fiber.Ctx) error {
	// Mendapatkan token dari header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Header otorisasi salah",
		})
	}

	tokenString := authHeader[7:]

	// Verifikasi token
	claims := new(models.JWTClaims)
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.JwtSecret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token tidak valid",
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Gagal mengautentikasi token",
		})
	}

	if !tkn.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token salah",
		})
	}

	// Menyimpan data user ke local context
	c.Locals("user", tkn)

	return c.Next()
}

func LogoutUser(c *fiber.Ctx) error {
	// Logout pada server hanya memberikan respons, frontend harus menghapus token dari penyimpanan lokal
	return c.JSON(fiber.Map{
		"message": "Logout berhasil, silakan hapus token di frontend",
	})
}
