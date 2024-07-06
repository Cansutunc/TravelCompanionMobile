import 'package:flutter/material.dart';

class ThemeService {
  Future<ThemeMode> themeMode() async => ThemeMode.light;
  Future<void> updateThemeMode(ThemeMode theme) async {}
}
