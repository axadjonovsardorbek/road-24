import 'dart:io';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../data/model/forms_status.dart';
import '../../data/model/image_model/image_model.dart';
import 'image_state.dart';

class ImageCubit extends Cubit<ImageState> {
  ImageCubit() : super(ImageState.initialValue());

  Future<void> insertImage({
    required ImageSource source,
    required ImageModel imageModel,
  }) async {
    emit(state.copyWith(status: FormsStatus.loading));
    final ImagePicker picker = ImagePicker();
    try {
      XFile? pickedFile = await picker.pickImage(source: source);
      if (pickedFile == null) {
        emit(
          state.copyWith(
            status: FormsStatus.pure,
            errorMessage: "Hech qanday rasm tanlanmadi.",
          ),
        );
      } else {
        Reference storageRef = FirebaseStorage.instance
            .ref()
            .child('file/images/${pickedFile.name}');
        File file = File(pickedFile.path);
        var uploadTask = await storageRef.putFile(file);
        String downloadUrl = await uploadTask.ref.getDownloadURL();
        ImageModel imagesModel = imageModel.copyWith(
          imagePath: downloadUrl,
        );
        emit(
          state.copyWith(
            status: FormsStatus.success,
            imagesModel: imagesModel,
          ),
        );
      }
    } catch (error) {
      emit(
        state.copyWith(
          status: FormsStatus.error,
          errorMessage: error.toString(),
        ),
      );
    }
  }

  Future<void> fetchAllImage() async {
    emit(
      state.copyWith(
        status: FormsStatus.fetch,
      ),
    );
    try {
      QuerySnapshot querySnapshot =
          await FirebaseFirestore.instance.collection("images").get();
      List<ImageModel> imageModel = querySnapshot.docs
          .map(
            (e) => ImageModel.fromJson(
              e.data() as Map<String, dynamic>,
            ),
          )
          .toList();
      emit(
        state.copyWith(status: FormsStatus.success, listImages: imageModel),
      );
    } catch (error) {
      emit(
        state.copyWith(
          status: FormsStatus.error,
          errorMessage: error.toString(),
        ),
      );
    }
  }

  void initialState() => emit(ImageState.initialValue());

  void deleteImage() => emit(
        state.copyWith(
          imagesModel: ImageModel.initialValue(),
        ),
      );

  void changeImage(ImageModel imageModel)=>state.copyWith(imagesModel: imageModel);
}
