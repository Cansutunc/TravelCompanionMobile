import 'package:carousel_slider/carousel_slider.dart';
import 'package:casu_mobile/pages/search_screen/bloc/pexel/pexel_images_cubit.dart';
import 'package:eva_icons_flutter/eva_icons_flutter.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';
import 'package:fluttericon/font_awesome5_icons.dart';
import 'package:shimmer/shimmer.dart';
import 'package:transparent_image/transparent_image.dart';

import '../../../common/colors.dart';
import '../../../common/global.dart';
import '../../../widgets/cast_widget_loader.dart';

class SearchListView extends StatelessWidget {
  const SearchListView({super.key, required this.movieRepository});

  final MovieRepository movieRepository;

  @override
  Widget build(BuildContext context) {
    // return PictureView(movieRepository: movieRepository);
    String location = Global.storageService.getLocationName();
    return Scaffold(
      body: BlocProvider(
        create: (_) => PexelImagesCubit(
          repository: context.read<MovieRepository>(),
        )..getPexelImage(location),
        child: PictureView(
          movieRepository: movieRepository,
        ),
      ),
    );
  }
}

class PictureView extends StatelessWidget {
  PictureView({super.key, required this.movieRepository});

  final MovieRepository movieRepository;
  final searchController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    final state = context.watch<PexelImagesCubit>().state;
    return Scaffold(
        appBar: AppBar(
          title: Row(
            children: [
             Expanded(child:  Container(
               height: 40.h,
               margin: EdgeInsets.only(top: 5.w),
               decoration: BoxDecoration(
                   color: Colors.white,
                   borderRadius: BorderRadius.all(Radius.circular(15.w)),
                   border: Border.all(color: Colors.black)),
               child: Row(
                 children: [
                   Container(
                     width: 16.w,
                     height: 16.w,
                     margin: EdgeInsets.only(left: 17.w),
                     child: Image.asset("assets/icons/search.png"),
                   ),
                   SizedBox(
                     width: 150.w,
                     height: 50.h,
                     child: TextField(
                       onChanged: (value) {

                       },
                       controller: searchController,
                       keyboardType: TextInputType.multiline,
                       decoration: InputDecoration(
                         hintText: "Search",
                         border: const OutlineInputBorder(
                             borderSide: BorderSide(color: Colors.transparent)),
                         enabledBorder: const OutlineInputBorder(
                             borderSide: BorderSide(color: Colors.transparent)),
                         disabledBorder: const OutlineInputBorder(
                             borderSide: BorderSide(color: Colors.transparent)),
                         focusedBorder: const OutlineInputBorder(
                             borderSide: BorderSide(color: Colors.transparent)),
                         hintStyle: TextStyle(
                           color: Colors.grey.withOpacity(0.5),
                         ),
                       ),
                       style: TextStyle(
                           color: AppColors.primaryText,
                           fontFamily: "Avenir",
                           fontWeight: FontWeight.normal,
                           fontSize: 14.sp),
                       autocorrect: false,
                     ),
                   )
                 ],
               ),

             )),
              Padding(
                padding: EdgeInsets.only(left: 10.dg),

                child: CircleAvatar(

                  radius: 20,
                  backgroundColor: Colors.grey,
                  child: IconButton(
                    icon: const Icon(
                      size: 30,
                      EvaIcons.search,
                      color: Colors.black,
                    ),
                    onPressed: () {
                      context.read<PexelImagesCubit>().getPexelImage(searchController.value.text);
                    },
                  ),
                ),
              ),
            ],
          ),
        ),
        body: ListView(children: [
          const Padding(
            padding: EdgeInsets.all(10.0),
          ),
          _body(context)
        ],)
    );
  }

  Widget _body(BuildContext context){
    final state = context.watch<PexelImagesCubit>().state;
    switch(state.status){
      case ListStatus.loading:
        return buildCastslistLoaderWidget(context);
      case ListStatus.success:

        return Stack(
          children: [
            CarouselSlider(
              options: CarouselOptions(
                autoPlay: false,
                viewportFraction: 1.0,
                aspectRatio: 2 / 2.8,
                enlargeCenterPage: true,
              ),
              items: state.images
                  .map((image) => Stack(
                children: [
                  Stack(
                    children: [
                      Shimmer.fromColors(
                        baseColor: Colors.black87,
                        highlightColor: Colors.white54,
                        enabled: true,
                        child: const AspectRatio(
                            aspectRatio: 2 / 2.8,
                            child: Icon(
                              FontAwesome5.film,
                              color: Colors.black26,
                              size: 40.0,
                            )),
                      ),
                      AspectRatio(
                          aspectRatio: 2 / 2.8,
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(5.0),
                            child: Image.network(image.imgUrl,fit: BoxFit.cover)
                          )),
                    ],
                  ),
                  AspectRatio(
                    aspectRatio: 2 / 2.8,
                    child: Container(
                      decoration: BoxDecoration(
                        gradient: LinearGradient(
                            begin: Alignment.bottomCenter,
                            end: Alignment.topCenter,
                            stops: const [
                              0.0,
                              0.4,
                              0.4,
                              1.0
                            ],
                            colors: [
                              Colors.black.withOpacity(1.0),
                              Colors.black.withOpacity(0.0),
                              Colors.black.withOpacity(0.0),
                              Colors.black.withOpacity(0.7),
                            ]),
                      ),
                    ),
                  ),
                  Positioned(
                      top: 15.0,
                      right: 10.0,
                      child: SafeArea(
                        child: Column(
                          children: [
                            const Text(
                              "Place: ",
                              style: TextStyle(
                                color: Colors.grey,
                                fontSize: 12.0,
                              ),
                            ),
                            Text(
                              image.title,
                              style: const TextStyle(
                                  fontWeight: FontWeight.bold,
                                  fontSize: 12.0,
                                  color: Colors.grey),
                            ),
                          ],
                        ),
                      )),
                ],
              ))
                  .toList(),
            ),
          ],
        );
      case ListStatus.failure:
      default:
        return buildCastslistLoaderWidget(context);
    }
  }

}
