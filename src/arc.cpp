#ifndef GRAPH_H_INCLUDED
#include "headers/graph.h"
#endif

template <typename T>
Graph<T>::Arc::Arc() {}

template <typename T>
Graph<T>::Arc::Arc (node_p theHead, node_p theTail, double theWeight) :
    head (theHead),
    tail (theTail),
    weight (theWeight) {}

template <typename T>
Graph<T>::Arc::Arc (node_p theHead, node_p theTail) :
    Graph<T>::Arc::Arc (theHead, theTail, 1) {}


template <typename T>
typename Graph<T>::node_p Graph<T>::Arc::getHead() {
    return head;
}

template <typename T>
typename Graph<T>::node_p Graph<T>::Arc::getTail() {
    return tail;
}

template <typename T>
double Graph<T>::Arc::getWeight() {
    return weight;
}

template <typename T>
void Graph<T>::Arc::setHead (node_p newHead) {
    head = newHead;
}

template <typename T>
void Graph<T>::Arc::setTail (node_p newTail) {
    tail = newTail;
}

template <typename T>
void Graph<T>::Arc::setWeight (double newWeight) {
    weight = newWeight;
}

template <typename T>
typename Graph<T>::Arc Graph<T>::Arc::copy() {
    return Graph<T>::Arc(head, weight);
}

template <typename T>
const std::string Graph<T>::Arc::toString() const {
    stringstream ss("");
    ss << head->getName();
    return ss.str();
}

template <typename T>
bool Graph<T>::Arc::equals(Graph<T>::Arc other) {
    // || for undirected?
    return head == other.getHead() && tail == other.getTail();
}

template <typename T>
const bool Graph<T>::Arc::isCoveredBy(const node_p node) const {
    return node == head || node == tail;
}
