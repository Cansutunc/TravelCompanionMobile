import 'dart:async';

import 'package:shared_preferences/shared_preferences.dart';

import 'constants.dart';


class StorageService {
  late final SharedPreferences _prefs;

  Future<StorageService> init() async {
    _prefs = await SharedPreferences.getInstance();
    return this;
  }

  Future<bool> setBool(String key, bool value) async {
    return await _prefs.setBool(key, value);
  }


  Future<bool> setString(String key, String value) async {
    return await _prefs.setString(key, value);
  }

  Future<bool> remove(String key){
    return _prefs.remove(key);
  }


  bool getDeviceFirstOpen() {
    return _prefs.getBool(AppConstants.STORAGE_DEVICE_OPEN_FIRST_TIME) ?? false;
  }


  bool getIsLoggedIn() {
    print(_prefs.getString(AppConstants.STORAGE_USER_TOKEN_KEY));
    return _prefs.getString(AppConstants.STORAGE_USER_TOKEN_KEY) == null
        ? false
        : true;
  }

  String getUsername(){
    var username = _prefs.getString(AppConstants.STORAGE_USER_PROFILE)??"";
    if (username.isNotEmpty){
      return username;
    }
    return "unknown";
  }

  String getLocationName(){
    var username = _prefs.getString(AppConstants.LOCATION)??"Istanbul";
    if (username.isNotEmpty){
      return username;
    }
    return "Istanbul";
  }

  String getUserId(){
    var userId = _prefs.getString(AppConstants.STORAGE_USER_TOKEN_KEY)??"";
    if (userId.isNotEmpty){
      return userId;
    }
    return "unknown";
  }


  String getProfilPath(){
    var profilPath = _prefs.getString(AppConstants.STORAGE_USER_PROFILE_PATH)??"";
    if (profilPath.isNotEmpty){
      return profilPath;
    }
    return "https://github.githubassets.com/assets/apple-touch-icon-144x144-b882e354c005.png";
  }
}
