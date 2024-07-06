import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../widgets/toast_info.dart';
import 'bloc/register_blocs.dart';

class RegisterController {
  final BuildContext context;

  const RegisterController({required this.context});

  Future<void> handleEmailRegister() async {
    // BlocProvider.of<SignInBloc>(context).state
    final state = context.read<RegisterBlocs>().state;
    String emailAddress = state.email;
    String password = state.password;
    String username = state.userName;
    String rePassword = state.rePassword;

    if (username.isEmpty) {
      toastInfo(msg: "You need to fill username");
    }

    if (emailAddress.isEmpty) {
      toastInfo(msg: "You need to fill email address");
    }

    if (password.isEmpty) {
      toastInfo(msg: "You need to fill password");
    }

    if (rePassword.isEmpty) {
      toastInfo(msg: "You need to fill rePassword");
    }

    try {
      final credential = await FirebaseAuth.instance
          .createUserWithEmailAndPassword(
              email: emailAddress, password: password);
      if (credential.user != null) {}
      await credential.user?.sendEmailVerification();
      await credential.user?.updateDisplayName(username);
      toastInfo(
          msg:
              "An email has been sent yo your registered email.To activate it please check your email box and click on the link");
      Navigator.of(context).pop();
    } on FirebaseAuthException catch (e) {
      if (e.code == "weak-password") {
        toastInfo(msg: "Thrown if the password is not strong enough");
      } else if (e.code == "email-already-in-use") {
        toastInfo(msg: "The email is already in use");
      } else if (e.code == "invalid-email") {
        toastInfo(msg: "Thrown if the email address is not valid");
      }
    }
  }
}
