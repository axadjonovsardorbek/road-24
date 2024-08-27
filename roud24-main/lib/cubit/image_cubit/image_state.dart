
import '../../data/model/forms_status.dart';
import '../../data/model/image_model/image_model.dart';

class ImageState {
  final String errorMessage;
  final FormsStatus status;
  final ImageModel imagesModel;
  final List<ImageModel> listImages;

  ImageState({
    required this.errorMessage,
    required this.status,
    required this.imagesModel,
    required this.listImages,
  });

  ImageState copyWith({
    String? errorMessage,
    FormsStatus? status,
    ImageModel? imagesModel,
    List<ImageModel>? listImages,
  }) =>
      ImageState(
        errorMessage: errorMessage ?? this.errorMessage,
        status: status ?? this.status,
        imagesModel: imagesModel ?? this.imagesModel,
        listImages: listImages ?? this.listImages,
      );

  static ImageState initialValue() => ImageState(
    errorMessage: '',
    status: FormsStatus.pure,
    imagesModel: ImageModel.initialValue(),
    listImages: [],
  );
}