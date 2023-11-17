echo 'find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r' >> lookagain.sh
find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r
