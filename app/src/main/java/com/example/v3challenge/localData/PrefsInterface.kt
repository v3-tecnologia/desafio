package com.example.v3challenge.localData

interface PrefsInterface {
	fun setPref(field: String)
	fun getPref(): String?
	fun removePref(field: String)
	fun clearPrefs()
}
