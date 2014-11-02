#!/usr/bin/env python
import graph_tool.all as gt
import optparse
import os
import re

import subprocess
from random import *
from joblib import Parallel, delayed

rangeRx = "(\d+)-(\d+),(\d+)"
options = None

def parseRange(m):
  return range(int(m.groups(0)[0]), int(m.groups(0)[1]) + 1, int(m.groups(0)[2]))

def main():
  p = optparse.OptionParser()
  p.add_option('--vertices', '-v', default="50,100,200,500")
  p.add_option('--degree_distribution', '-d', default="50,100,200,500")
  p.add_option('--filename', '-f')
  p.add_option('--output_format', '-o', default="jpg")
  p.add_option('--layout', '-l', default="sfdp")
  p.add_option('--show', '-s', dest="show", action="store_true")
  global options
  options, arguments = p.parse_args()
  mv =re.match(rangeRx, options.vertices)
  md = re.match(rangeRx, options.degree_distribution)
  if hasattr(mv, 'groups'):
    if hasattr(md,'groups'):
      for verts in parseRange(mv):
        for deg in parseRange(md):
          work(verts, deg)
    else:
        for verts in parseRange(mv):
          for deg in options.degree_distribution.split(','):
            work(verts, deg)
  else:
    if hasattr(md,'groups'):
      for verts in options.vertices.split(','):
        for deg in parseRange(md):
          work(verts, deg)
    else:
      for verts in options.vertices.split(','):
        for deg in options.degree_distribution.split(','):
          work(verts, deg)

def sample_k(max):
  accept = False
  while not accept:
    k = randint(1,max+1)
    accept = random() < 1.0/k
  return k

def work(verts, deg):
  if ~hasattr(options, 'filename'):
    options.filename = str(verts) + '_' + str(deg)
  print(options.filename)
  g = gt.random_graph(int(verts), lambda: sample_k(int(deg)), model="probabilistic",
             vertex_corr=lambda i, k: 1.0 / (1 + abs(i - k)), directed=False,
             n_iter=100)
  dotfile = os.getcwd() + '/'+options.filename+'.dot'
  imgfile = os.getcwd() + '/'+options.filename+'.'+options.output_format
  g.save(dotfile, 'dot')
  subprocess.Popen([options.layout, "-T", options.output_format], stdin=open(dotfile, 'r'), stdout=open(imgfile, 'a')).communicate()
  if options.show:
    subprocess.Popen(['display', imgfile])

if __name__ == '__main__':
  main()
