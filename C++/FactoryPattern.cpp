#include<iostream>
using namespace std;

enum class ShapeType {
	DEFAULT,
	CIRCLE,
	SQUARE
};

class figureFeatures {
public:
	virtual void showShape() = 0;
	virtual ~figureFeatures() = default; // to avoid memory leak, since c++ doesn't have an inbuilt garbage collector
};

class Circle: public figureFeatures {
public:
	void showShape() override {
		cout << "Got a Circle\n";
	}
};


class Square : public figureFeatures {
public:
	void showShape() override {
		cout << "Got a Square\n";
	}
};

figureFeatures* getFactoryImplementation(ShapeType requiredShape) {
	switch (requiredShape) {
	case ShapeType::CIRCLE:
		return new Circle();
	case ShapeType::SQUARE:
		return new Square();
	default:
		return NULL;
	}
}

int main() {
	string s;
	cin >> s;

	ShapeType requiredShape = ShapeType::DEFAULT;

	if (s == "Circle") {
		requiredShape = ShapeType::CIRCLE;
	}
	else if (s == "Square") {
		requiredShape = ShapeType::SQUARE;
	}

	auto relevantImplementation = getFactoryImplementation(requiredShape);
	if (relevantImplementation != NULL) relevantImplementation->showShape();
	else cout << "DAMN\n";

	delete(relevantImplementation);
}
