plugins {
    id("com.android.application")
    id("org.jetbrains.kotlin.android")
}

android {
    namespace = "br.com.camargoguilherme.desafio_android_embarcado"
    compileSdk = 34

    defaultConfig {
        applicationId = "br.com.camargoguilherme.desafio_android_embarcado"
        minSdk = 25
        targetSdk = 33
        versionCode = 2
        versionName = "1.1"

        testInstrumentationRunner = "androidx.test.runner.AndroidJUnitRunner"
    }


    buildTypes {
        debug {
            isDebuggable = true
            isMinifyEnabled = false
        }
        release {
            isMinifyEnabled = true
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


    implementation("androidx.appcompat:appcompat:1.7.0")
    implementation("com.google.android.material:material:1.12.0")
    implementation("com.google.android.gms:play-services-location:21.3.0")
    implementation("androidx.test:core-ktx:1.6.1")
    implementation("com.squareup.okhttp3:okhttp:4.12.0")

    // JUnit para testes unitários
    testImplementation("junit:junit:4.13.2")

    // Mockk para criar mocks e spyers
    testImplementation("io.mockk:mockk:1.13.4")
    testImplementation("com.squareup.okhttp3:mockwebserver:4.12.0")

    // Robolectric para rodar testes unitários de componentes Android sem precisar de um dispositivo físico
    testImplementation("org.robolectric:robolectric:4.6.1")

    // AndroidX Test para rodar testes em Android
    androidTestImplementation("androidx.test.ext:junit:1.2.1")
    androidTestImplementation("androidx.test:core:1.6.1")
    androidTestImplementation("androidx.core:core-ktx:1.9.0")
    androidTestImplementation("androidx.test.espresso:espresso-core:3.6.1")







}