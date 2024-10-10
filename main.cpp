a#include <iostream>
#include <string>
#include <list>

using namespace std;

// Предварительное объявление классов
//class Client;
//class User;
//class Detail;
//class Order;

// Класс User
class User {
protected:
    struct Admin {
        string username;
        string pass;
    };

private:
    string name;
    string surname;
    string email;
    string position;

public:
    bool registration() {
        // Логика регистрации
        cout << "Регистрация пользователя\n";
        return true;
    }

    bool authorization() {
        // Логика авторизации
        cout << "Авторизация пользователя\n";
        return true;
    }

    bool addOrder() {
        // Логика добавления заказа
        cout << "Добавление заказа\n";
        return true;
    }

    bool removeOrder() {
        // Логика удаления заказа
        cout << "Удаление заказа\n";
        return true;
    }

    bool editOrder() {
        // Логика редактирования заказа
        cout << "Редактирование заказа\n";
        return true;
    }

    bool addClient() {
        // Логика добавления клиента
        cout << "Добавление клиента\n";
        return true;
    }

    bool removeClient() {
        // Логика удаления клиента
        cout << "Удаление клиента\n";
        return true;
    }

    bool editClient() {
        // Логика редактирования клиента
        cout << "Редактирование клиента\n";
        return true;
    }

    bool addDetail() {
        // Логика добавления детали
        cout << "Добавление детали\n";
        return true;
    }

    bool removeDetail() {
        // Логика удаления детали
        cout << "Удаление детали\n";
        return true;
    }

    bool editDetail() {
        // Логика редактирования детали
        cout << "Редактирование детали\n";
        return true;
    }
};

// Класс Program
class Program {
private:
    //list<Client> clients;
    list<User> admins;
    //list<Detail> details;
    //list<Order> orders;
    string date;

public:
    bool reboot() {
        // Логика перезагрузки программы
        cout << "Перезагрузка программы\n";
        return true;
    }

    void showOrders() {
        // Логика вывода заказов
        cout << "Показ заказов\n";
    }

    void showClients() {
        // Логика вывода клиентов
        cout << "Показ клиентов\n";
    }

    void showAdmins() {
        // Логика вывода администраторов
        cout << "Показ администраторов\n";
    }

    void showDetails() {
        // Логика вывода деталей
        cout << "Показ деталей\n";
    }
};

// Пример использования
int main() {
    User user;
    Program program;

    user.registration();
    user.authorization();
    user.addOrder();
    user.removeOrder();
    user.editOrder();
    user.addClient();
    user.removeClient();
    user.editClient();
    user.addDetail();
    user.removeDetail();
    user.editDetail();

    program.reboot();
    program.showOrders();
    program.showClients();
    program.showAdmins();
    program.showDetails();

    return 0;
}
