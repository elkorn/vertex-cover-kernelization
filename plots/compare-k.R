setwd("/home/elkorn/Code/go/src/github.com/elkorn/vertex-cover-kernelization/main/measurement-results")

cr <- read.table("MeasureKernelizationCrownReduction", header=TRUE, sep="\t")
crp <- read.table("MeasureKernelizationCrownReductionPreprocessing", header=TRUE, sep="\t")
maxnum <- function(col) {
  return(max(sapply(col, as.numeric)))
}

lwd = 2

yMax = max(maxnum(cr$Ts), maxnum(crp$Ts))
xMax = max(length(cr$Ts))
x <- seq(100, 2000, 100)
c <- rainbow(4)

plot(x, x, xlab = "Liczba wierzchołków grafu", ylab="Czas redukcji dziedziny [s]", main="Wpływ przetwarzania wstępnego na czas redukcji koron", ylim=c(0, yMax), pch = 20, type="n")
mkLineX <- function(x, data, idx) {
  lines(x, data$Ts, col=c[idx], lwd = lwd, type="o")
}

mkLine <- function(data, idx) {
  n = length(data$Ts)
  print(n)
  mkLineX(x[1:n], data, idx)
}


mkLine(cr, 1)
mkLine(crp, 2)
legend(100,yMax, c("Bez przetwarzania wstępnego", "Z przetwarzaniem wstępnym"), lty = c(1,1,1,1), lwd=c(lwd,lwd,lwd,lwd),col=c)