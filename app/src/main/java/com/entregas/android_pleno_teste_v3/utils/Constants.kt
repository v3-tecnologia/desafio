package com.entregas.android_pleno_teste_v3.utils

class Constants private constructor() {

    companion object {
        const val REQUEST_LOCATION_PERMISSION = 1
        const val REQUEST_CAMERA_PERMISSION = 2
        const val REQUEST_STORAGE_PERMISSION = 3

        const val CHANNEL_ID = "BackgroundServiceChannel"
        const val NOTIFICATION_ID = 1
        const val GYROSCOPE_DATA_ACTION = "com.entregas.android_pleno_teste_v3.GYROSCOPE_DATA"
        const val LOCATION_DATA_ACTION = "com.entregas.android_pleno_teste_v3.LOCATION_DATA"

        const val EXTRA_X = "x"
        const val EXTRA_Y = "y"
        const val EXTRA_Z = "z"
        const val EXTRA_TIMESTAMP = "timestamp"
        const val EXTRA_LATITUDE = "latitude"
        const val EXTRA_LONGITUDE = "longitude"

        const val ACTION_REQUEST_PERMISSIONS = "com.entregas.android.REQUEST_PERMISSIONS"
        const val ACTION_BOOT_COMPLETED = "android.intent.action.BOOT_COMPLETED"

        fun getRequestCode(requestType: String): Int {
            return when (requestType) {
                "location" -> REQUEST_LOCATION_PERMISSION
                "camera" -> REQUEST_CAMERA_PERMISSION
                "storage" -> REQUEST_STORAGE_PERMISSION
                else -> -1
            }
        }
    }
}
