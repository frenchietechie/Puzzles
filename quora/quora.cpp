#include <iostream>
#include <cstring>
#include <list>

class Point;

int width;
int height;
Point* start = NULL ;
int zeroCount;
//2D matrix of point*
Point*** datacenter = NULL ;
int solutionsCount = 0;

class Point {
public:
	bool visited;
	//number of ways remaining to reach that point (0 to 4)
	int reachCount;
	int i, j;
	int value;
	std::list<Point*> neighbours;
	Point(int i, int j, int value);
	void Init();
	void Visit(int visited);
};

Point::Point(int i, int j, int value) {
	this-> i = i;
	this-> j = j;
	this->value = value;
	this->reachCount = 0;
	this->visited = false;
}

void Point::Init() {
	if (value == 2)
		start = this;
	if (value == 0)
		++zeroCount;
	if (value == 1)
		return ;

	if (i + 1 < width && datacenter[j][i + 1]->value != 1) {
		neighbours.push_front(datacenter[j][i + 1]);
		++reachCount;
	}
	if (i - 1 >= 0 && datacenter[j][i - 1]->value != 1) {
		neighbours.push_front(datacenter[j][i - 1]);
		++reachCount;
	}
	if (j + 1 < height && datacenter[j + 1][i]->value != 1) {
		neighbours.push_front(datacenter[j + 1][i]);
		++reachCount;
	}
	if (j - 1 >= 0 && datacenter[j - 1][i]->value != 1) {
		neighbours.push_front(datacenter[j - 1][i]);
		++reachCount;
	}
}

void Point::Visit(int visited) {
	if (value == 3) {
		if (visited == zeroCount)
			++solutionsCount;

		return;
	} else {
		this->visited = true;
		std::list<Point*>::iterator it;
		for (it = neighbours.begin(); it != neighbours.end(); ++it) {
			Point* neighbour = (*it);
			neighbour->reachCount--;
		}
		//marker to see if a visit has been forced
		bool hadTo = false;
		for (it = neighbours.begin(); it != neighbours.end(); ++it) {
			Point* neighbour = (*it);
			if ((!neighbour->visited) && (neighbour->reachCount <= 1) && (neighbour->value != 3)) {
				hadTo = true;
				neighbour->Visit(visited + 1);
				break;
			}
		}
		if (!hadTo) {
			for (it = neighbours.begin(); it != neighbours.end(); ++it) {
				Point* neighbour = (*it);
				if ((!neighbour->visited)) {
					neighbour->Visit(visited + 1);
				}
			}
		}
		this->visited = false;
		for (it = neighbours.begin(); it != neighbours.end(); ++it) {
			Point* neighbour = (*it);
			neighbour->reachCount++;
		}
	}
}

void parseInput() {
	std::cin >> width;
	std::cin >> height;

	datacenter = new Point**[height];
	for (int j = 0; j < height; j++) {
		datacenter[j] = new Point*[width];
		for (int i = 0; i < width; i++) {
			int data;
			std::cin >> data;
			datacenter[j][i] = new Point(i, j, data);
		}
	}

	for (int j = 0; j < height; j++)
		for (int i = 0; i < width; i++)
			datacenter[j][i]->Init();

}

int main() {
	parseInput();
	start->Visit(-1);
	std::cout << solutionsCount << std::endl;
	return 0;
}
