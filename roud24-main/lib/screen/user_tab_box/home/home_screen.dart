import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:roud24/cubit/image_cubit/image_cubit.dart';
import 'package:roud24/utils/extension/extension.dart';

import '../../../utils/color/app_colors.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Мой автомобиль'),
        leading: IconButton(
          icon: const Icon(Icons.person),
          onPressed: () {},
        ),
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              _buildCarInfo(context),
              const SizedBox(height: 16),
              _buildFinesSection(),
              const SizedBox(height: 16),
              _buildDetailsSection(),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildCarInfo(BuildContext context) {
    return Card(
      elevation: 4,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Row(
          children: [
            CachedNetworkImage(
              width: 100.w,
              height: 80,
              fit: BoxFit.cover,
              imageUrl:
                  context.read<ImageCubit>().state.imagesModel.imagePath.isEmpty
                      ? context.read<ImageCubit>().state.listImages[0].imagePath
                      : context.read<ImageCubit>().state.imagesModel.imagePath,
              placeholder: (context, url) => const CupertinoActivityIndicator(
                color: AppColors.c000000,
              ),
            ),
            const SizedBox(width: 16),
            const Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Мой автомобиль',
                  style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
                ),
                Text('Chevrolet Spark'),
                SizedBox(height: 4),
                Text('01 U 123 DB'),
              ],
            ),
            const Spacer(),
            IconButton(
              icon: const Icon(Icons.edit),
              onPressed: () {},
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildFinesSection() {
    return Card(
      elevation: 4,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      child: const Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Штрафы',
              style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
            ),
            SizedBox(height: 8),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  '788 000 сум',
                  style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                ),
                Text(
                  'скидка 30%',
                  style: TextStyle(color: Colors.green),
                ),
              ],
            ),
            SizedBox(height: 4),
            Text(
              '2 штрафа не оплачено',
              style: TextStyle(color: Colors.red),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildDetailsSection() {
    return Column(
      children: [
        Row(
          children: [
            _buildDetailCard('Страховка', '123 дня', 'до 12 авг 2021'),
            const SizedBox(width: 16),
            _buildDetailCard('Тех. осмотр', '98 дней', 'до 17 апр 2021'),
          ],
        ),
        const SizedBox(height: 16),
        Row(
          children: [
            _buildDetailCard('Доверенность', '236 дней', ''),
            const SizedBox(width: 16),
            _buildDetailCard('Тонировка', '15 дней', ''),
          ],
        ),
      ],
    );
  }

  Widget _buildDetailCard(String title, String duration, String subtitle) {
    return Expanded(
      child: Card(
        elevation: 4,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                title,
                style:
                    const TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
              const SizedBox(height: 8),
              Text(
                duration,
                style:
                    const TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
              ),
              const SizedBox(height: 4),
              Text(subtitle),
            ],
          ),
        ),
      ),
    );
  }
}
