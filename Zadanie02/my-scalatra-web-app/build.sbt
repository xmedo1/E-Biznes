val ScalatraVersion = "3.1.2"

ThisBuild / scalaVersion := "3.3.7"
ThisBuild / organization := "com.example"

lazy val hello = (project in file("."))
  .settings(
    name := "My Scalatra Web App",
    version := "0.1.0-SNAPSHOT",
    libraryDependencies ++= Seq(
      "org.scalatra" %% "scalatra-jakarta" % ScalatraVersion,
      "org.scalatra" %% "scalatra-scalatest-jakarta" % ScalatraVersion % "test",
      "ch.qos.logback" % "logback-classic" % "1.5.19" % "runtime",
    ),
  )

enablePlugins(SbtTwirl)
enablePlugins(SbtWar)

Test / fork := true
