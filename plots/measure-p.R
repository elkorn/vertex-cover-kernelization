setwd("/home/elkorn/Code/go/src/github.com/elkorn/vertex-cover-kernelization/main/measurement-results")

p <- read.table("MeasurePreprocessing", header=TRUE, sep="\t")
maxnum <- function(col) {
  return(max(sapply(col, as.numeric)))
}

lwd = 2

yMax = max(maxnum(p$Ts))
xMax = max(length(cr$Ts))
x <- seq(100, 2000, 100)
c <- rainbow(4)

plot(p$V, p$Ts, xlab = "Liczba wierzchołków grafu", ylab="Czas [s]", main="Czas przetwarzania wstępnego", ylim=c(0, yMax), pch = 20, type="n")
mkLineX <- function(x, data, idx) {
  lines(x, data$Ts, col=c[idx], lwd = lwd, type="o")
}

mkLine <- function(data, idx) {
  n = length(data$Ts)
  print(n)
  mkLineX(x[1:n], data, idx)
}

mkLine(p, 1)