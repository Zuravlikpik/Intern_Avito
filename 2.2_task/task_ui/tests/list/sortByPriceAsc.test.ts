import { test, expect } from "@playwright/test";
import { ListPage } from "../../pages/listPage/listPage";

test("Главная страница: сортировка по цене по возрастанию", async ({ page }) => {
    const listPage = new ListPage(page);

    await listPage.open();

    // сортировка по возрастанию
    await listPage.sortSelect.selectOption("price");
    await listPage.orderSelect.selectOption("asc");
    await listPage.page.waitForTimeout(1000);

    let hasNext = true;
    let pageCount = 0;

    while (hasNext && pageCount < 3) {
        const prices = await listPage.getPrices();
        for (let i = 0; i < prices.length - 1; i++) {
            expect(prices[i]).toBeLessThanOrEqual(prices[i + 1]);
        }

        hasNext = await listPage.goToNextPageIfExists();
        pageCount++;
    }
});
