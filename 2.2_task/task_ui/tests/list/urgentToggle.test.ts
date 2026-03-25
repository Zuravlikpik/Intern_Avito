import { test, expect } from "@playwright/test";
import { ListPage } from "../../pages/listPage/listPage";

test("Главная страница: фильтр 'Только срочные'", async ({ page }) => {
    const listPage = new ListPage(page);

    await listPage.open();

    //включ.тогл
    await listPage.enableUrgentOnly();

    await listPage.assertPriorityUrgent();

    let hasNext = true;
    let pageCount = 0;

    while (hasNext && pageCount < 3) {
        const total = await listPage.getCardsCount();
        const urgent = await listPage.getUrgentCount();

        expect(total).toBeGreaterThan(0);
        expect(urgent).toBe(total);

        hasNext = await listPage.goToNextPageIfExists();
        pageCount++;
    }
});
