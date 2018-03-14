SELECT w1.id FROM Weather w1 JOIN Weather w2 ON TO_DAYS(w1.date) = TO_DAYS(w2.date) + 1 AND w1.Temperature > w2.Temperature;
