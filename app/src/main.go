package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func main() {
    // Inicializa o viper para ler o arquivo de configuração
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println("Error reading config file, ", err)
    }

    // Inicializa o gin para criar rotas
    r := gin.Default()

    // Cria a rota de healthcheck
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "OK",
        })
    })

    // Inicia a API com as configurações lidas do arquivo de configuração
    r.Run(":" + viper.GetString("server.port"))
}