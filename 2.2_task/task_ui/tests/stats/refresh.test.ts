import { test, expect } from "@playwright/test";
import { StatsPage } from "../../pages/statsPage/statsPage";

test("Статистика: кнопка Обновить работает после остановки таймера", async ({ page }) => {
    const statsPage = new StatsPage(page);

    await statsPage.open();

    await statsPage.stopTimer();

    // после стопа таймера нет  и жмём обновить
    await statsPage.refreshButton.click();

    // таймер должен появиться снова
    await expect(statsPage.timeValue).toBeVisible();
});
