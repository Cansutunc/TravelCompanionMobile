import 'package:flutter/material.dart';
import 'package:casu_mobile/common/theme/theme_controller.dart';
import 'package:casu_mobile/common/theme/theme_service.dart';
import 'app.dart';
import 'common/global.dart';

Future<void> main() async {
  await Global.init();
  final themeController = ThemeController(ThemeService());
  await themeController.loadSettings();
  runApp(App(
    themeController: themeController,
  ));
}
