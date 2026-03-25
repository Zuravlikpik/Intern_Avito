import { test, expect } from "@playwright/test";
import { StatsPage } from "../../pages/statsPage/statsPage";

test("Статистика: кнопка Обновить сбрасывает таймер", async ({ page }) => {
    const statsPage = new StatsPage(page);

    await statsPage.open();

    const before = await statsPage.getTimerText();

    await statsPage.refreshAndWaitForReset(before);

    const after = await statsPage.getTimerText();

    const beforeSec = statsPage.parseTimeToSeconds(before);
    const afterSec = statsPage.parseTimeToSeconds(after);

    // таймер стал больше (откатился)
    expect(afterSec).toBeGreaterThan(beforeSec);

    await expect(statsPage.refreshButton).toBeEnabled();
});
