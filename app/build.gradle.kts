plugins {
    alias(libs.plugins.android.application)
    alias(libs.plugins.jetbrains.kotlin.android)
    id("org.jetbrains.kotlin.kapt")
    }

android {
    namespace = "com.entregas.android_pleno_teste_v3"
    compileSdk = 34

    defaultConfig {
        applicationId = "com.entregas.android_pleno_teste_v3"
        minSdk = 24
        targetSdk = 34
        versionCode = 1
        versionName = "1.0"

        testInstrumentationRunner = "androidx.test.runner.AndroidJUnitRunner"
    }

    buildTypes {
        release {
            isMinifyEnabled = false
            proguardFiles(
                getDefaultProguardFile("proguard-android-optimize.txt"),
                "proguard-rules.pro"
            )
        }
    }
    compileOptions {
        sourceCompatibility = JavaVersion.VERSION_1_8
        targetCompatibility = JavaVersion.VERSION_1_8
    }
    kotlinOptions {
        jvmTarget = "1.8"
    }
}

dependencies {
    // Core libraries
    implementation(libs.androidx.core.ktx)
    implementation(libs.androidx.appcompat)
    implementation(libs.material)
    implementation(libs.androidx.activity)
    implementation(libs.androidx.constraintlayout)

    // New dependencies
    implementation(libs.coroutines.core) // kotlinx-coroutines-core
    implementation(libs.coroutines.android) // kotlinx-coroutines-android
    implementation(libs.retrofit) // retrofit
    implementation(libs.okhttp) // okhttp
    implementation(libs.koin.core) // koin-core
    implementation(libs.koin.android) // koin-android

    implementation( "androidx.room:room-runtime:2.6.1")
    kapt("androidx.room:room-compiler:2.6.1")
    implementation ("androidx.room:room-ktx:2.2.1")
    // Test dependencies
    testImplementation(libs.junit)
    testImplementation(libs.mockito.core)
    testImplementation(libs.mockk)
    testImplementation(libs.truth)
    testImplementation(libs.robolectric)
    androidTestImplementation(libs.androidx.junit)
    androidTestImplementation(libs.androidx.espresso.core)
}
