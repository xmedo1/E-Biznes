package com.example

import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.application.*
import io.ktor.server.netty.*
import kotlinx.coroutines.launch
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json

@Serializable
data class DiscordMessage(val content: String)

fun main(args: Array<String>): Unit = EngineMain.main(args)

fun Application.module() {
    configureSerialization()
    configureRouting()
    val webhookUrl = environment.config.property("discord.webhook_url").getString()

    val client = HttpClient(CIO) {
        install(io.ktor.client.plugins.contentnegotiation.ContentNegotiation) {
            json(Json { ignoreUnknownKeys = true })
        }
    }

    launch {
        try {
            client.post(webhookUrl) {
                contentType(ContentType.Application.Json)
                setBody(DiscordMessage("Wiadomosc testowa"))
            }
            println("Wyslano wiadomosc")
        } catch (e: Exception) {
            println("Blad podczas wysylania: ${e.message}")
        }
    }
}