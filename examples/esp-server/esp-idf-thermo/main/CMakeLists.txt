idf_component_register(
    SRCS "main.c" "emi/cjson/cJSON.c"
    INCLUDE_DIRS "." "emi"
    REQUIRES esp_http_server esp_wifi esp_event esp_netif nvs_flash esp_timer
)

# No libcurl on the ESP32: the generated server code only needs the url helpers.
target_compile_definitions(${COMPONENT_LIB} PRIVATE FETCHX_NO_CURL)
