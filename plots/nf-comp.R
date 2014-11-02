setwd("/home/elkorn/Code/go/src/github.com/elkorn/vertex-cover-kernelization/main/measurement-results")

nf <- read.table("results_vc_nf", header=TRUE, sep="\t")
nfp <- read.table("results_vc_nf_p", header=TRUE, sep="\t")
maxnum <- function(col) {
  return(max(sapply(col, as.numeric)))
}

lwd = 2


yMax = max(maxnum(nf$Ts), maxnum(nfp$Ts))
x <- seq(100, 2000, 100)
c <- rainbow(3)
plot(x, x, xlab = "Liczba wierzchołków grafu", ylab="Czas wyznaczania pokrycia wierzchołkowego [s]", main="Wyznaczanie pokrycia wierzchołkowego w oparciu o przepływ w sieci", ylim=c(0, yMax), pch = 20, type="n")
mkLine <- function(data, idx) {
  lines(x, data$Ts, col=c[idx], lwd = lwd, type="o")
}

mkLineX <- function(x, data, idx) {
  lines(x, data$Ts, col=c[idx], lwd = lwd, type="o")
}

mkLineX(x[1:17], nf, 1)
mkLineX(x, nfp, 2)
legend(100,yMax, c("Bez przetwarzania wstępnego","Z przetwarzaniem wstępnym"), lty = c(1,1), lwd=c(lwd,lwd),col=c)       


