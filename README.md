## Overall Workflow

1. Install tools:
```bash
./install_gps_tools.sh
```

2. Get metadata:
```bash
./extended_metadata_extractor.sh Sample-PNG-Image.png
```

3. Update metadata:
```bash
./add_custom_metadata.sh Sample-PNG-Image.png
```

4. Get updated metadata:
```bash
./extended_metadata_extractor.sh Sample-PNG-Image.png
```

## install_gps_tools.sh

Usage:
```bash
./install_gps_tools.sh
```

## add_custom_metadata.sh

Usage:
```bash
./add_custom_metadata.sh <image_file>
```

## extended_metadata_extractor.sh

Usage:
```bash
./extended_metadata_extractor.sh <image_file>
```

## gps_metadata_extractor.sh

Usage:
```bash
./gps_metadata_extractor.sh <image_file>
```
```
