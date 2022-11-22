FROM tensorflow/serving:2.9.2

COPY monitoring_config.txt /monitoring_config.txt
COPY batching_config.txt /batching_config.txt

COPY testdata/saved_model_half_plus_two_cpu /models/half_plus_two
ENV MODEL_NAME half_plus_two


