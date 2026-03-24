import { test, expect } from "@playwright/test";
import { ListPage } from "../../pages/listPage/listPage";

test("Фильтр по диапазону цен", async ({ page }) => {
    const listPage = new ListPage(page);

    // открываем страницу объявлений
    await listPage.open();

    // получаем текущие цены на первой странице
    const prices = await listPage.getPrices();
    expect(prices.length).toBeGreaterThan(0);

    // рассчитываем диапазон цен
    const minPrice = Math.min(...prices);
    const maxPrice = Math.max(...prices);

    // фильтруем средний диапазон
    const from = Math.floor(minPrice + (maxPrice - minPrice) * 0.25);
    const to = Math.ceil(minPrice + (maxPrice - minPrice) * 0.75);

    console.log(`Диапазон цен: от ${from} до ${to}`);

    // Применяем фильтр диапазона цен
    await listPage.applyPriceRangeFilter(from, to);

    // Проверяем все страницы
    let hasNext = true;
    while (hasNext) {
        const filteredPrices = await listPage.getPrices();
        for (const price of filteredPrices) {
            expect(price).toBeGreaterThanOrEqual(from);
            expect(price).toBeLessThanOrEqual(to);
        }

        hasNext = await listPage.goToNextPageIfExists();
    }
});
