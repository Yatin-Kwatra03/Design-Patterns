#include <iostream>
#include <queue>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <chrono>

std::queue<int> q;
std::mutex mtx;
std::condition_variable cv;

void producer() {
    for (int i = 1; i <= 5; ++i) {
        {
            std::unique_lock<std::mutex> lock(mtx);
            q.push(i);
            std::cout << "Produced: " << i << std::endl;
        } // lock released here

        cv.notify_one(); // wake up one waiting consumer
        std::this_thread::sleep_for(std::chrono::milliseconds(500)); // simulate work
    }
}

void consumer() {
    for (int i = 1; i <= 5; ++i) {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, [] { return !q.empty(); }); // wait until queue is not empty

        int item = q.front();
        q.pop();
        std::cout << "Consumed: " << item << std::endl;
    }
}

int main() {
    std::thread t1(producer);
    std::thread t2(consumer);

    t1.join();
    t2.join();

    return 0;
}
