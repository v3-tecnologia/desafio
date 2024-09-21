package com.example.v3challenge.network

sealed class Resource<T>(
    val data: T? = null,
    val message: String? = null,
    val code: Int? = null,
    val stringRes: Int? = null,
    val showError: Boolean = true,
    val exception: Exception? = null
) {
    class Success<T>(data: T) : Resource<T>(data)
    class Error<T>(
        message: String,
        code: Int? = null,
        stringRes: Int? = null,
        showError: Boolean = true,
        exception: Exception? = null,
        data: T? = null
    ) :
        Resource<T>(data, message, code, stringRes, showError, exception)
}