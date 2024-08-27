import 'package:easy_localization/easy_localization.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:roud24/servicec/firebase_options.dart';
import 'package:roud24/servicec/services_locator.dart';

import 'app/app.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await EasyLocalization.ensureInitialized();
  Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);
  setUp();
  runApp(
    EasyLocalization(
      supportedLocales: const [
        Locale("en", "EN"),
        Locale("ru", "RU"),
        Locale("uz", "UZ"),
      ],
      path: 'assets/translation',
      fallbackLocale: const Locale("uz","UZ"),
      child: const App(),
    ),
  );
}