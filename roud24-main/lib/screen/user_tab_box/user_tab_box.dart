import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:roud24/screen/user_tab_box/home/home_screen.dart';
import 'package:roud24/screen/user_tab_box/info/info_screen.dart';

import '../../cubit/tab_box_cubit/tab_box_cubit.dart';
import '../../utils/color/app_colors.dart';

class TabBoxScreen extends StatefulWidget {
  const TabBoxScreen({super.key});

  @override
  State<TabBoxScreen> createState() => _TabBoxScreenState();
}

class _TabBoxScreenState extends State<TabBoxScreen> {
  final GlobalKey<CurvedNavigationBarState> _bottomNavigationKey = GlobalKey();
  List<Widget> screens = [
    const HomeScreen(),
    const InfoScreen(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocBuilder<TabBoxCubit, int>(
        builder: (context, state) {
          return IndexedStack(
            index: state,
            children: screens,
          );
        },
      ),
      bottomNavigationBar: BlocBuilder<TabBoxCubit, int>(
        builder: (context, state) {
          return CurvedNavigationBar(
            key: _bottomNavigationKey,
            index: state,
            items: const <Widget>[
              Icon(
                Icons.home,
                size: 24,
              ),
              Icon(
                Icons.info_outline,
                size: 24,
              )
            ],
            backgroundColor: Colors.blueAccent,
            animationCurve: Curves.easeInOut,
            animationDuration: const Duration(milliseconds: 600),
            onTap: (index) {
              context.read<TabBoxCubit>().changeIndex(index);
              //_showExplanation(index);
            },
            letIndexChange: (index) => true,
          );
        },
      ),
    );
  }
}
