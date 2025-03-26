FROM python:3.12-slim AS builder

# Install and configure Poetry
ENV POETRY_VERSION=1.7.1
#RUN curl -sSL https://install.python-poetry.org/ | python -
#ENV PATH="${PATH}:/root/.local/bin"
RUN pip install poetry==${POETRY_VERSION}

RUN poetry config virtualenvs.create false

# Copy dependencies
COPY ./poetry.lock ./pyproject.toml /app/
# Set /app/ as working directory
WORKDIR /app/

# Install dependencies
RUN poetry install --no-interaction --no-ansi --compile --without=lint,dev,seeding --no-root

# Copy source files
COPY ./src /app/src

RUN poetry install --no-interaction --no-ansi --compile --only-root

FROM python:3.12-slim AS final

ENV PYTHONUNBUFFERED=1

COPY --from=builder /usr/local/lib/python3.12/site-packages /usr/local/lib/python3.12/site-packages
COPY --from=builder /app/ /app/

WORKDIR /app/

ENTRYPOINT ["python", "src/timeseries/main.py"]