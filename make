#!/usr/bin/env bash

rm $(find templates -type f -name "*.go")

cat > templates/templates.go <<\EOF
package templates

var Data = map[string] func() []byte {
EOF

for tmpl in $(find templates -type f); do
	out=$(echo $tmpl | sed -r 's/templates\///g' | sed -r 's/\//z/g' | sed -r 's/\./zz/g')
	file=$(echo $tmpl | sed -r 's/templates\///g')

	if [[ $out != "templateszzgo" ]]; then
		echo "Building $file ...";
		go-bindata -pkg="templates" -out="templates/$out.go" -func="$out" $tmpl;
		echo -e "\t\"$file\": $out," >> templates/templates.go;
	fi
done

echo "}" >> templates/templates.go
go fmt templates/*.go
