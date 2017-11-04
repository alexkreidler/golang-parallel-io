# The example/benchmark file names
files=(specs recursive parallel walk)

# Build each example/benchmark
for file in ${files[@]}; do
  go build $file.go
done

# Set the directory to test on
DIR=~ #$(pwd)

# Initialize the commands with `find` as a baseline
cmd=("find $DIR")

# Add the Go examples
for file in ${files[@]}; do
  # Don't run specs as a benchmark
  if [ "$file" == "specs" ]; then
    continue
  fi
  # cmd=("${cmd[@]}" "./$file $DIR")
  cmd+=("./$file $DIR")
done

# List CPU cores/parallel specs
./specs

# Run each command, timing it and counting the lines output
# See https://stackoverflow.com/questions/9084257/bash-array-with-spaces-in-elements
for ((i = 0; i < ${#cmd[@]}; i++))
do
    c="${cmd[$i]}"
    echo $c
    time $c | wc -l
done
