#include<iostream>
using namespace std;


class Button {
public:
	virtual void click() = 0;
	virtual ~Button() = default;
};


class CheckBox {
public:
	virtual void toggle() = 0;
	virtual ~CheckBox() = default;
};



class WindowsButton: public Button {
public:
	void click() override {
		cout << "windows button clicked\n";
	}
};


class WindowsCheckBox: public CheckBox {
public:
	void toggle() override {
		cout << "windows CheckBox toggled\n";
	}
};

class MacButton: public Button {
public:
	void click() override {
		cout << "Mac button clicked\n";
	}
};


class MacCheckBox: public CheckBox {
public:
	void toggle() override {
		cout << "Mac CheckBox toggled\n";
	}
};


class GUIFactory {
public:
	virtual Button* createButton() = 0;
	virtual CheckBox* createCheckBox() = 0;
	~GUIFactory() = default;
};

// Here we didn't import Button and Checkbox, windows factory is not a button
// or checkbox, it's a GUIFactory. Understand while import what u import is what u are.
class WindowsFactory: public GUIFactory {
public:
	Button* createButton() {
		return new WindowsButton();
	}

	CheckBox* createCheckBox() {
		return new WindowsCheckBox();
	}
};

class MacFactory: public GUIFactory {
public:
	Button* createButton() {
		return new MacButton();
	}

	CheckBox* createCheckBox() {
		return new MacCheckBox();
	}
};

void renderUI(GUIFactory* factory) {
	auto button = factory->createButton();
	auto checkBox = factory->createCheckBox();

	button->click();
	checkBox->toggle();

	delete(button);
	delete(checkBox);
	delete(factory);
}

int main() {
	string os = "Mac";

	GUIFactory* factory;

	if (os == "Windows") {
		factory = new WindowsFactory();
	}
	else {
		factory = new MacFactory();
	}

	renderUI(factory);
}

