package br.com.camargoguilherme.desafio_android_embarcado

import android.content.Context
import android.os.Environment
import android.util.Log
import java.io.File
import java.io.FileWriter
import java.io.IOException
import java.text.SimpleDateFormat
import java.util.Date
import java.util.Locale

class LogWriter(private val context: Context) {

    // Função para gerar o nome do arquivo com a data
    private fun generateLogFileName(): String {
        val dateFormat = SimpleDateFormat("yyyyMMdd", Locale.getDefault())
        val date = Date()
        return "log_${dateFormat.format(date)}.txt"
    }

    // Função para escrever uma mensagem no arquivo de log
    fun writeLog(tag: String, message: String) {
        try {
            // Diretório de Downloads
            val downloadsDir = context.getExternalFilesDir(Environment.DIRECTORY_DOWNLOADS)

            if (downloadsDir != null && downloadsDir.exists()) {
                val logFileName = generateLogFileName()
                val logFile = File(downloadsDir, logFileName)
                val writer = FileWriter(logFile, true) // true para habilitar o modo de append

                val dateFormat = SimpleDateFormat("yyyy-MM-dd HH:mm:ss", Locale.getDefault())
                val date = Date()

                writer.append("${dateFormat.format(date)} ${tag} - ${message}")
                writer.append("\n")
                writer.flush()
                writer.close()

                Log.i(tag, message)
                Log.e(tag, message)
                Log.d(tag, message)
                Log.v(tag, message)

                Log.v(tag, "Log escrito com sucesso em ${logFile.absolutePath}")
            } else {
                Log.e(tag, "Erro ao acessar o diretório Downloads")
            }
        } catch (e: IOException) {
            Log.e(tag, "Erro ao escrever log: ${e.message}")
        }
    }
}
