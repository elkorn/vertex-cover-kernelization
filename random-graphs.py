#!/usr/bin/env python
import graph_tool.all as gt
import optparse
import os

from subprocess import call
from random import *

def sample_k(max):
  accept = False
  while not accept:
    k = randint(1,max+1)
    accept = random() < 1.0/k
  return k

def main():
  p = optparse.OptionParser()
  p.add_option('--vertices', '-v', default=30)
  p.add_option('--degree_distribution', '-d', default=30)
  p.add_option('--filename', '-f', default="test")
  p.add_option('--output_format', '-o', default="jpg")
  p.add_option('--layout', '-l', default="sfdp")

  options, arguments = p.parse_args()
  g = gt.random_graph(int(options.vertices), lambda: sample_k(int(options.degree_distribution)), model="probabilistic",
             vertex_corr=lambda i, k: 1.0 / (1 + abs(i - k)), directed=False,
             n_iter=100)
  dotfile = os.getcwd() + '/'+options.filename+'.dot'
  g.save(dotfile, 'dot')

  with open(dotfile, 'r') as fin:
    print(fin.read())

if __name__ == '__main__':
  main()
