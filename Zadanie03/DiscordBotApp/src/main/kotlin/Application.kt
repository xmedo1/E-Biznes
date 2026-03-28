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
import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent

@Serializable
data class DiscordMessage(val content: String)
val categoriesData = mapOf(
    "Warzywa" to listOf("Ziemniak", "Marchewka", "Salata"),
    "Owoce" to listOf("Banan", "Jablko", "Pomarancza"),
)
fun main(args: Array<String>): Unit = EngineMain.main(args)

fun Application.module() {
    configureSerialization()
    configureRouting()
    val webhookUrl = environment.config.property("discord.webhook_url").getString()
    val token = environment.config.property("discord.token").getString()

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

    launch {
        val kord = Kord(token)
        kord.on<MessageCreateEvent> {
            val content = message.content
            println("Wiadomosc od: ${message.author?.username} | Tresc: $content")
            when {
                content.lowercase() == "test" -> {
                    message.channel.createMessage("dziala!!!")
                }

                content.lowercase() == "kategorie" -> {
                    val categories = categoriesData.keys.joinTo(StringBuilder(), separator = ", ", prefix = "[", postfix = "]")
                    message.channel.createMessage("Kategorie: ${categories}")
                }
            }

        }
        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }
}