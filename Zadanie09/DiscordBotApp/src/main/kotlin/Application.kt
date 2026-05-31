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
import io.ktor.server.routing.*
import io.ktor.server.http.content.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.client.call.*

@Serializable
data class DiscordMessage(val content: String)

val categoriesData = mapOf(
    "warzywa" to listOf("Ziemniak", "Marchewka", "Salata"),
    "owoce" to listOf("Banan", "Jablko", "Pomarancza"),
)

@Serializable
data class FrontendRequest(val text: String)

@Serializable
data class FrontendResponse(val reply: String)

@Serializable
data class PythonRequest(val text: String)

@Serializable
data class PythonResponse(val response: String)

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

    routing {
        staticResources("/", "static")
        post("/api/chat") {
            val userMsg = call.receive<FrontendRequest>()

            try {
                val pyResponse = client.post("http://localhost:5000/askai") {
                    contentType(ContentType.Application.Json)
                    setBody(PythonRequest(text = userMsg.text))
                }.body<PythonResponse>()

                call.respond(FrontendResponse(reply = pyResponse.response))
            } catch (e: Exception) {
                call.respond(FrontendResponse(reply = "Blad polaczenia z AI."))
            }
        }
    }

    launch {
        try {
            val pyResponse = client.post("http://localhost:5000/askai") {
                contentType(ContentType.Application.Json)
                setBody(PythonRequest(text = "[start]"))
            }.body<PythonResponse>()

            client.post(webhookUrl) {
                contentType(ContentType.Application.Json)
                setBody(DiscordMessage("${pyResponse.response} (Zacznij wiadomość od `!ai`, żeby ze mną porozmawiać! :)"))
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

            if (message.author?.isBot == true) return@on

            when {
                content.lowercase() == "test" -> {
                    message.channel.createMessage("dziala!!!")
                }

                content.lowercase() == "kategorie" -> {
                    val categories =
                        categoriesData.keys.joinTo(StringBuilder(), separator = ", ", prefix = "[", postfix = "]")
                    message.channel.createMessage("Kategorie: ${categories}")
                }

                categoriesData.containsKey(content.lowercase()) -> {
                    val products = categoriesData[content.lowercase()]?.joinTo(
                        StringBuilder(),
                        separator = ", ",
                        prefix = "[",
                        postfix = "]"
                    )
                    message.channel.createMessage("Produkty w kategorii ${content.lowercase()}: $products")
                }

                content.lowercase().startsWith("!ai ") -> {
                    val question = content.drop(4).trim()
                    try {
                        val pyResponse = client.post("http://localhost:5000/askai") {
                            contentType(ContentType.Application.Json)
                            setBody(PythonRequest(text = question))
                        }.body<PythonResponse>()

                        // discord message limit
                        val safeResponse = if (pyResponse.response.length > 1900) {
                            pyResponse.response.take(1900) + "..."
                        } else {
                            pyResponse.response
                        }

                        message.channel.createMessage("🤖 AI: $safeResponse")
                    } catch (e: Exception) {
                        println("Error: ${e.message}")
                        e.printStackTrace()
                        message.channel.createMessage("🤖 AI: coś chyba nie działa ajajaj")
                    }
                }
            }
        }

        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }
}