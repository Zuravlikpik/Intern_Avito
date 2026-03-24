import { test } from "@playwright/test";
import { ListPage } from "../../pages/listPage/listPage";

test("Главная страница: сортировка по цене по убыванию", async ({ page }) => {
    const listPage = new ListPage(page);

    await listPage.open();
    await listPage.sortByPriceDesc();

    let hasNext = true;
    let pageCount = 0;

    while (hasNext && pageCount < 3) {
        const prices = await listPage.getPrices();

        await listPage.assertSortedDesc(prices);

        hasNext = await listPage.goToNextPageIfExists();
        pageCount++;
    }
});
