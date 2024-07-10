import pandas as pd
import numpy as np
from sklearn.metrics.pairwise import cosine_similarity

# Step 1: Load data
data = {'user_id': [1, 1, 1, 2, 2, 3, 3, 3],
        'item_id': [1, 2, 3, 1, 4, 2, 3, 5],
        'rating': [5, 3, 4, 4, 5, 2, 5, 4]}

df = pd.DataFrame(data)

# Step 2: Create user-item matrix
user_item_matrix = df.pivot_table(index='user_id', columns='item_id', values='rating').fillna(0)

# Step 3: Compute user similarity
user_similarity = cosine_similarity(user_item_matrix)

# Step 4: Make recommendations
def recommend(user_id, user_item_matrix, user_similarity, top_n=3):
    user_index = user_id - 1
    similarity_scores = user_similarity[user_index]
    user_ratings = user_item_matrix.iloc[user_index]

    # Predict ratings
    predicted_ratings = np.dot(similarity_scores, user_item_matrix) / np.array([np.abs(similarity_scores).sum(axis=0)])

    # Recommend top N items the user hasn't rated yet
    recommendations = pd.Series(predicted_ratings, index=user_item_matrix.columns)
    recommendations = recommendations[user_ratings == 0]
    return recommendations.nlargest(top_n)

# Example: Recommend items for user 1
recommendations = recommend(1, user_item_matrix, user_similarity)
print("Recommended items for user 1:", recommendations.index.tolist())

