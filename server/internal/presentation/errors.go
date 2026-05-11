package presentation

import "github.com/gin-gonic/gin"

var (
	DNSNotFound = gin.H{
		"error": gin.H{
			"code":    "DNS_NOT_FOUND",
			"message": "dns not found in query",
		},
	}

	DNSInvalid = gin.H{
		"error": gin.H{
			"code":    "DNS_INVALID",
			"message": "dns address is not a valid IP",
		},
	}

	DNSAlreadyExists = gin.H{
		"error": gin.H{
			"code":    "DNS_ALREADY_EXISTS",
			"message": "dns already exists",
		},
	}

	FileNotFound = gin.H{
		"error": gin.H{
			"code":    "FILE_NOT_FOUND",
			"message": "dns file not found",
		},
	}

	InternalError = gin.H{
		"error": gin.H{
			"code":    "INTERNAL_ERROR",
			"message": "internal error",
		},
	}
)
