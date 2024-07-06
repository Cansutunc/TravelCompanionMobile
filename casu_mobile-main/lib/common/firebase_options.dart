// ignore_for_file: type=lint
import 'package:firebase_core/firebase_core.dart' show FirebaseOptions;
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, kIsWeb, TargetPlatform;

class DefaultFirebaseOptions {
  static FirebaseOptions get currentPlatform {
    if (kIsWeb) {
      return web;
    }
    switch (defaultTargetPlatform) {
      case TargetPlatform.android:
        return android;
      case TargetPlatform.iOS:
        return ios;
      case TargetPlatform.macOS:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for macos - '
              'you can reconfigure this by running the FlutterFire CLI again.',
        );
      case TargetPlatform.windows:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for windows - '
              'you can reconfigure this by running the FlutterFire CLI again.',
        );
      case TargetPlatform.linux:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for linux - '
              'you can reconfigure this by running the FlutterFire CLI again.',
        );
      default:
        throw UnsupportedError(
          'DefaultFirebaseOptions are not supported for this platform.',
        );
    }
  }

  static const FirebaseOptions web = FirebaseOptions(
    apiKey: "AIzaSyDS90eIPN8j-Hcb-Pv83NuxvBORAAm8YnE",
    authDomain: "casu2-506a7.firebaseapp.com",
    projectId: "casu2-506a7",
    storageBucket: "casu2-506a7.appspot.com",
    messagingSenderId: "520940893635",
    appId: "1:520940893635:web:f9f15a7a1cf7d6cc343976"
  );

  static const FirebaseOptions android = FirebaseOptions(
    apiKey: 'AIzaSyBwomO2R7dCIw0jLSI8oadO7MwVBkJ4FhY',
    appId: '1:625002506307:android:e47f09b75555e7b4250cf6',
    messagingSenderId: '520940893635',
    projectId: 'casu2-506a7',
    storageBucket: 'casu2-506a7.casu2.com',
  );

  static const FirebaseOptions ios = FirebaseOptions(
    apiKey: 'AIzaSyAaZNBOTJZMqtcT7iT3V4gupEEAx-FyR7c',
    appId: '1:625002506307:ios:66eaae940592e8b8250cf6',
    messagingSenderId: '520940893635',
    projectId: 'casu2-506a7',
    storageBucket: 'casu2-506a7.casu2.com',
    iosBundleId: 'com.example.chatApp',
  );
}