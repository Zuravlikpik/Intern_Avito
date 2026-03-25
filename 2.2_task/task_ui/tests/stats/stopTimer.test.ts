import { test } from "@playwright/test";
import { StatsPage } from "../../pages/statsPage/statsPage";

test("Статистика: остановка таймера отключает автообновление", async ({ page }) => {
    const statsPage = new StatsPage(page);

    await statsPage.open();

    await statsPage.stopTimer();
});
