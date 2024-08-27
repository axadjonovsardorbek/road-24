import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:roud24/cubit/image_cubit/image_cubit.dart';
import 'package:roud24/cubit/tab_box_cubit/tab_box_cubit.dart';
import 'package:roud24/screen/route.dart';
import 'package:roud24/screen/splash/splash_screen.dart';

import '../utils/color/app_colors.dart';
import '../utils/extension/extension.dart';

class App extends StatefulWidget {
  const App({super.key});

  @override
  State<App> createState() => _AppState();
}

class _AppState extends State<App> {
  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (_) => ImageCubit()..fetchAllImage(),
        ),
        BlocProvider(
          create: (_) => TabBoxCubit(),
        ),
      ],
      child: const MyApp(),
    );
  }
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    width = context.getWidth();
    height = context.getHeight();
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        appBarTheme: AppBarTheme(
          systemOverlayStyle: SystemUiOverlayStyle(
            statusBarIconBrightness: Brightness.light,
            statusBarBrightness: Brightness.light,
            statusBarColor: AppColors.c356899.withOpacity(0.6),
          ),
        )
      ),
      locale: context.locale,
      localizationsDelegates: context.localizationDelegates,
      supportedLocales: context.supportedLocales,
      home: const SplashScreen(),
      initialRoute: RouteName.splash,
      onGenerateRoute: AppRoute.generateRoute,
    );
  }
}
