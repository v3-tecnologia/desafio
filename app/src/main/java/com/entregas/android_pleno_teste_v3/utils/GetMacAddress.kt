package com.entregas.android_pleno_teste_v3.utils

import android.content.Context
import android.net.wifi.WifiInfo
import android.net.wifi.WifiManager
import android.provider.Settings

class GetMacAddress(val context: Context) {

    private fun getMacAddressFromWifiManager(): String? {
        val wifiManager = context.getSystemService(Context.WIFI_SERVICE) as WifiManager
        val wifiInfo: WifiInfo = wifiManager.connectionInfo
        return wifiInfo.macAddress
    }

    fun getUniqueDeviceId(): String {
        return Settings.Secure.getString(context.contentResolver, Settings.Secure.ANDROID_ID)
    }
}