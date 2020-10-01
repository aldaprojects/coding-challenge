// Edad en Días
// https://www.urionlinejudge.com.br/judge/es/problems/view/1020

#include <bits/stdc++.h>

using namespace std;

int main()
{
    int initialDays;
    cin >> initialDays;

    int years = initialDays/365;
    int months = initialDays%365/30;
    int days = initialDays%365%30;

    printf("%d ano(s)\n%d mes(es)\n%d dia(s)", years, months, days);

    //cout << years << "ano(s)" << endl;
    //cout << months << "mes(es)" << endl;
    //cout << days << "dia(s)" << endl;

    return 0;
}
