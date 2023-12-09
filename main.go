package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalculationRequest struct {
	X int `json:"x"` // Capacity of jug X
	Y int `json:"y"` // Capacity of jug Y
	Z int `json:"z"` // Target volume in jug Z
}

type JugState struct {
	x, y, z int // Current volume in each jug
}

type Step struct {
	Description string   `json:"description"`
	State       JugState `json:"state"`
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func calculateHandler(c *gin.Context) {
	var req CalculationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure Z is divisible by the smaller of X or Y
	smallerJug := min(req.X, req.Y)
	if req.Z%smallerJug != 0 {
		c.JSON(http.StatusOK, gin.H{"result": "No Solution"})
		return
	}

	// Initialize the state
	var steps []Step
	currentState := JugState{0, 0, 0}

	// Attempt to solve the problem
	for currentState.z != req.Z {
		// Fill the smaller jug and transfer to the larger one
		if smallerJug == req.X && currentState.x < req.X {
			currentState.x = req.X
			steps = append(steps, Step{"Fill X", currentState})
		} else if smallerJug == req.Y && currentState.y < req.Y {
			currentState.y = req.Y
			steps = append(steps, Step{"Fill Y", currentState})
		}

		// Transfer from the smaller to the larger jug
		if smallerJug == req.X && currentState.y < req.Y {
			transfer := min(currentState.x, req.Y-currentState.y)
			currentState.x -= transfer
			currentState.y += transfer
			steps = append(steps, Step{"Transfer X to Y", currentState})
		} else if smallerJug == req.Y && currentState.x < req.X {
			transfer := min(currentState.y, req.X-currentState.x)
			currentState.y -= transfer
			currentState.x += transfer
			steps = append(steps, Step{"Transfer Y to X", currentState})
		}

		// Transfer from the larger jug to Z
		if currentState.z < req.Z {
			transfer := min(max(currentState.x, currentState.y), req.Z-currentState.z)
			if currentState.x == max(currentState.x, currentState.y) {
				currentState.x -= transfer
				steps = append(steps, Step{"Transfer X to Z", currentState})
			} else {
				currentState.y -= transfer
				steps = append(steps, Step{"Transfer Y to Z", currentState})
			}
			currentState.z += transfer
		}
	}

	c.JSON(http.StatusOK, steps)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/api/calculate", calculateHandler)
	router.Run(":8080")
}