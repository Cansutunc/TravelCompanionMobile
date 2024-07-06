import 'dart:io';

import 'package:casu_mobile/model/cast_response.dart';
import 'package:casu_mobile/model/genre_response.dart';
import 'package:casu_mobile/model/location_response.dart';
import 'package:casu_mobile/model/movie_detail_response.dart';
import 'package:casu_mobile/model/movie_response.dart';
import 'package:casu_mobile/model/pexel_image_response.dart';
import 'package:casu_mobile/model/user_location.dart';
import 'package:casu_mobile/model/user_location_response.dart';
import 'package:dio/dio.dart';
import 'package:casu_mobile/model/person_response.dart';
import 'package:casu_mobile/model/video_response.dart';
import 'package:geolocator/geolocator.dart';

import '../model/hobbie.dart';
import '../model/hobbie_response.dart';
import '../model/route.dart';
import '../model/route_response.dart';

class MovieRepository {
  final String apiKey = "8a1227b5735a7322c4a43a461953d4ff";
  static String mainUrl = "https://api.themoviedb.org/3";
  final Dio _dio = Dio();

  var getUpComingApi = '$mainUrl/movie/upcoming';
  var getPopularMoviesApi = '$mainUrl/movie/popular';
  var getTopRatedMoviesApi = '$mainUrl/movie/top_rated';
  var getNowPlayingMoviesApi = '$mainUrl/movie/now_playing';

  var getMoviesApi = '$mainUrl/movie/top_rated';
  var getMoviesUrl = '$mainUrl/discover/movie';
  var getPlayingUrl = '$mainUrl/movie/now_playing';
  var getGenresUrl = "$mainUrl/genre/movie/list";
  var getPersonsUrl = "$mainUrl/trending/person/week";
  var movieUrl = "$mainUrl/movie";
  var getPexelImagesUrl = "https://api.pexels.com/v1/search";
  var getHobbiesUrl = "http://192.168.1.109:5002/v1/book";
  var getRoutesUrl = "http://192.168.1.109:5002/v1/city";
  var getAcceptedRoutesUrl = "http://192.168.1.109:5002/v1/city/acceptAll";
  var getUserLocationUrl = "http://192.168.1.109:5002/v1/country";
  var removeHobbiesUrl = "http://192.168.1.109:5002/v1/dispatch/book/book-remove";
  var removeRoutesUrl = "http://192.168.1.109:5002/v1/dispatch/city/city-remove";
  var acceptRoutesUrl = "http://192.168.1.109:5002/v1/city/accept";
  var deleteRoutesUrl = "http://192.168.1.109:5002/v1/city/remove";
  var addHobbiesUrl = "http://192.168.1.109:5002/v1/dispatch/book/book-create";
  var addRoutesUrl = "http://192.168.1.109:5002/v1/dispatch/city/city-create";
  var addUserLocationUrl = "http://192.168.1.109:5002/v1/dispatch/country/country-create";
  var getLocationUrl = "https://geocode.maps.co/reverse";


  Future<PexelImageResponse> getPexelImages(String query) async {
    var params = {"query": query, "per_page": 5, "page": 1};
    try {
      _dio.options.headers["authorization"] = 'DTRUAd8bF0bOB6uYPFs7yQfLhRBFDLyRXKzRErPc2kOJmied1E3lkXUk';
      Response response =
      await _dio.get(getPexelImagesUrl, queryParameters: params);
      return PexelImageResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return PexelImageResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }



  Future<HobbieResponse> getHobbies() async {
    try {
      Response response =
      await _dio.get(getHobbiesUrl);
      return HobbieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return HobbieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<RouteResponse> getRoutes() async {
    try {
      Response response =
      await _dio.get(getRoutesUrl);
      return RouteResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return RouteResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<RouteResponse> getAcceptedRoutes(String username) async {
    try {
      var params = {"username": username};
      Response response =
      await _dio.get(getAcceptedRoutesUrl,queryParameters: params);
      return RouteResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return RouteResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<UserLocationResponse> getUserLocations() async {
    try {
      Response response =
      await _dio.get(getUserLocationUrl);
      return UserLocationResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return UserLocationResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }


  Future<HobbieResponse> removeHobbies(String id) async {
    print(id);
    try {
      Response response =
      await _dio.post(removeHobbiesUrl,data: {"id":id});
      return HobbieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return HobbieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<RouteResponse> removeRoute(String id) async {

    try {
      Response response =
      await _dio.post(removeRoutesUrl,data: {"id":id});
      return RouteResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return RouteResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<bool> acceptRoute(String id,String username) async {
    try {
      var params = {"id": id, "username": username};
      Response response =
      await _dio.get(acceptRoutesUrl,queryParameters: params);
      if(response.statusCode == HttpStatus.ok){
        return true;
      }
      return false;
    } catch (error, stacktrace) {
      return false;
    }
  }

  Future<bool> deleteRoute(String id,String username) async {
    try {
      var params = {"id": id, "username": username};
      Response response =
      await _dio.get(deleteRoutesUrl,queryParameters: params);
      if(response.statusCode == HttpStatus.ok){
        return true;
      }
      return false;
    } catch (error, stacktrace) {
      return false;
    }
  }

  Future<bool> addHobbie(Hobbie hobbie) async {
    try {
      Response response =
      await _dio.post(addHobbiesUrl,data: hobbie.toMap() );
      if(response.statusCode == HttpStatus.ok){
        return true;
      }
      return false;

    } catch (error, stacktrace) {
      return false;
    }
  }

  Future<bool> addRoute(Routes route) async {
    try {
      Response response =
      await _dio.post(addRoutesUrl,data: route.toMap() );
      if(response.statusCode == HttpStatus.ok){
        return true;
      }
      return false;

    } catch (error, stacktrace) {
      return false;
    }
  }

  Future<bool> addUserLocation(UserLocation userLocation) async {
    try {
      Response response =
      await _dio.post(addUserLocationUrl,data: userLocation.toMap() );
      if(response.statusCode == HttpStatus.ok){
        return true;
      }
      return false;

    } catch (error, stacktrace) {
      return false;
    }
  }

  Future<LocationResponse> getLocation(Position position) async {
    var params = {"api_key": "66047447dc1a0394257330zji8d70fa", "lon": position.longitude, "lat": position.latitude};
    try {
      Response response =
      await _dio.get(getLocationUrl, queryParameters: params);
      return LocationResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return LocationResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }



  Future<MovieResponse> getMovies(int page) async {
    var params = {"api_key": apiKey, "language": "en-US", "page": page};
    try {
      Response response = await _dio.get(getMoviesApi, queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieResponse> getPopularMovies(int page) async {
    var params = {"api_key": apiKey, "language": "en-US", "page": page};
    try {
      Response response =
          await _dio.get(getPopularMoviesApi, queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieResponse> getNowPlaying(int page) async {
    var params = {"api_key": apiKey, "language": "en-US", "page": page};
    try {
      Response response =
          await _dio.get(getNowPlayingMoviesApi, queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieResponse> getUpcomingMovies() async {
    var params = {"api_key": apiKey, "language": "en-US", "page": 1};
    try {
      Response response =
          await _dio.get(getUpComingApi, queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieResponse> getTopRatedMovies() async {
    var params = {"api_key": apiKey, "language": "en-US", "page": 1};
    try {
      Response response =
          await _dio.get(getTopRatedMoviesApi, queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieResponse> getPlayingMovies() async {
    var params = {"api_key": apiKey, "language": "en-US", "page": 1};
    try {
      Response response =
          await _dio.get(getPlayingUrl, queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<GenreResponse> getGenres() async {
    var params = {"api_key": apiKey, "language": "en-US"};
    try {
      Response response = await _dio.get(getGenresUrl, queryParameters: params);
      return GenreResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return GenreResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<PersonResponse> getPersons() async {
    var params = {"api_key": apiKey};
    try {
      Response response =
          await _dio.get(getPersonsUrl, queryParameters: params);
      return PersonResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return PersonResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieResponse> getMovieByGenre(int id) async {
    var params = {
      "api_key": apiKey,
      "language": "en-US",
      "page": 1,
      "with_genres": id
    };
    try {
      Response response = await _dio.get(getMoviesUrl, queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieDetailResponse> getMovieDetail(int id) async {
    var params = {"api_key": apiKey, "language": "en-US"};
    try {
      Response response =
          await _dio.get(movieUrl + "/$id", queryParameters: params);
      return MovieDetailResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieDetailResponse.withError(
          "Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<VideoResponse> getMovieVideos(int id) async {
    var params = {"api_key": apiKey, "language": "en-US"};
    try {
      Response response = await _dio.get(movieUrl + "/$id" + "/videos",
          queryParameters: params);
      return VideoResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return VideoResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<MovieResponse> getSimilarMovies(int id) async {
    var params = {"api_key": apiKey, "language": "en-US"};
    try {
      Response response = await _dio.get(movieUrl + "/$id" + "/similar",
          queryParameters: params);
      return MovieResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return MovieResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }

  Future<CastResponse> getCasts(int id) async {
    var params = {"api_key": apiKey, "language": "en-US"};
    try {
      Response response = await _dio.get(movieUrl + "/$id" + "/credits",
          queryParameters: params);
      return CastResponse.fromJson(response.data);
    } catch (error, stacktrace) {
      return CastResponse.withError("Error: $error, StackTrace: $stacktrace");
    }
  }
}
