args <- commandArgs(trailing = TRUE)
fn <- args[1]
df <- read.table(fn, sep = ",", header = TRUE)
print(summary(df))
input_file <- args[1]
file <- file.info(input_file)$size

if (file > 1500000000){
    z <- input_file
    z
    print(paste(z, "is over file limit"))
} else {
    z <- input_file
    print(paste("[ * ] Byte size -> ",file))
    print(paste(z, "is fine to read....."))
    Data = read.csv(input_file)
    print(Data)
}

