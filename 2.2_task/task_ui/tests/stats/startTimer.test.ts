import { test, expect } from "@playwright/test";
import { StatsPage } from "../../pages/statsPage/statsPage";

test("Статистика: повторный запуск таймера продолжает отсчет", async ({ page }) => {
    const statsPage = new StatsPage(page);

    await statsPage.open();

    const beforeStop = await statsPage.getTimerText();
    const beforeSec = statsPage.parseTimeToSeconds(beforeStop);

    await statsPage.stopTimer();

    await statsPage.startTimer();

    const afterStart = await statsPage.getTimerText();
    const afterSec = statsPage.parseTimeToSeconds(afterStart);

    // не должен сброситься на 5:00
    expect(afterStart).not.toBe("5:00");

    // должен быть примерно тем же
    expect(afterSec).toBeLessThanOrEqual(beforeSec);
});
