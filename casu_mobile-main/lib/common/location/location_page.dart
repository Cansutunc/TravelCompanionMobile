import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
// import 'package:frontend/core/base/location_cubit/route_cubit.dart';
import 'package:geolocator/geolocator.dart';




// class LocationPage extends StatelessWidget {
//   final String? _currentAddress = null;
//   final Position? _currentPosition = null;
//   const LocationPage({super.key});
//   @override
//   Widget build(BuildContext context) {
//     return Center(
//       child: SafeArea(
//         child: Center(
//           child: Column(
//             mainAxisAlignment: MainAxisAlignment.center,
//             children: [
//               Text('LAT: ${_currentPosition?.latitude ?? ""}'),
//               Text('LNG: ${_currentPosition?.longitude ?? ""}'),
//               Text('ADDRESS: ${_currentAddress ?? ""}'),
//               Text(BlocProvider.of<LocationCubit>(context).state.latitude.toString()),
//               const SizedBox(height: 32),
//               ElevatedButton(
//                 onPressed: () async {
//                   context.read<LocationCubit>().getLocation();
//                 },
//                 child: const Text("Get Current Location"),
//               ),
//             ],
//           ),
//         ),
//       ),
//     );
//   }
// }
