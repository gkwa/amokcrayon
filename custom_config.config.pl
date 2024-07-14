%Image::ExifTool::UserDefined = (
    'Image::ExifTool::XMP::Main' => {
        eden => {
            SubDirectory => {
                TagTable => 'Image::ExifTool::UserDefined::eden',
            },
        },
    },
);

%Image::ExifTool::UserDefined::eden = (
    GROUPS => { 0 => 'XMP', 1 => 'XMP-eden', 2 => 'Image' },
    NAMESPACE => { 'eden' => 'http://ns.example.com/eden/1.0/' },
    WRITABLE => 'string',
    Product => {},
    Brand => {},
    Weight => {},
    Sheets => {},
    Price => {},
    IsGlutenFree => {},
    IsVegan => {},
);

1;

